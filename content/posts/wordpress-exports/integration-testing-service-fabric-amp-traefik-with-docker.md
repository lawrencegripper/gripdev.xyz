---
author: gripdev
category:
  - docker
date: "2017-12-22T14:27:23+00:00"
guid: http://blog.gripdev.xyz/?p=1121
summary: |-
  Here is the plan:

  1. Use docker to run a Service Fabric Linux cluster in a container
  2. Deploy a test app to the cluster and create 25 instances of it

  Aim: While developing the Traefik SF integration it will provide a simple cluster to use, debug and perform integration testing.

  \*TLDR: Have a look the full code [in this PR](https://github.com/containous/traefik-extra-service-fabric/pull/21)

  It was a fun journey but I got it working...
tag:
  - docker
  - servicefabric
  - testing
  - traefik
title: Integration testing Service Fabric &amp; Traefik with Docker
url: /2017/12/22/integration-testing-service-fabric-traefik-with-docker/

---
Here is the plan:

1. Use docker to run a Service Fabric Linux cluster in a container
1. Deploy a test app to the cluster and create 25 instances of it

Aim: While developing the Traefik SF integration it will provide a simple cluster to use, debug and perform integration testing.

\*TLDR: Have a look the full code [in this PR](https://github.com/containous/traefik-extra-service-fabric/pull/21)

It was a fun journey but I got it working...

## First hurdle: Get a SF Cluster image which starts quickly.

Our starting point is this [guide here](https://docs.microsoft.com/en-us/azure/service-fabric/service-fabric-get-started-mac). In it you pull an the ServiceFabric docker image and then execute `./setup.sh` and `./run.sh`, which are scripts contained in the image.

**Note**: You need to add the additional docker daemon config to run SF in a container. See [details here.](https://docs.microsoft.com/en-us/azure/service-fabric/service-fabric-get-started-mac#create-a-local-container-and-set-up-service-fabric)

Once done you have a cluster running in the container. Unfortunately `./setup.sh` takes a while to run, installing packages from apt, such as the JVM, and then setting up some environment configuration.

Fortunately, we can build our own image on top of this to makes things quicker. Essentially, we’ll run the `./setup.sh` as part in our docker build so it doesn't have to run when we start the container.

My first attempt at this had a `dockerfile` like so:

\[code lang=text\]
FROM servicefabricoss/service-fabric-onebox
WORKDIR /home/ClusterDeployer
RUN ./setup.sh
EXPOSE 19080 19000 80 443
CMD ./run.sh
\[/code\]

Now if you build an image from this it will start and you can connect to the explorer. However, when you try and upload to the ImageStore you’ll get a `500 Code` response.

So after some messing around I took at look at the `setup.sh` script and found this at the end:

\[code lang=bash\]
/etc/init.d/ssh start
locale-gen en\_US.UTF-8
export LANG=en\_US.UTF-8
export LANGUAGE=en\_US:en
export LC\_ALL=en\_US.UTF-8

\[/code\]

These actions wouldn’t be capture in the docker image we built as they’re not persisted. So we can update our Dockerfile to capture these like so:

\[code lang=text\]
FROM servicefabricoss/service-fabric-onebox
WORKDIR /home/ClusterDeployer
RUN ./setup.sh
#Generate the local
RUN locale-gen en\_US.UTF-8
#Set environment variables
ENV LANG=en\_US.UTF-8
ENV LANGUAGE=en\_US:en
ENV LC\_ALL=en\_US.UTF-8
EXPOSE 19080 19000 80 443
#Start SSH before running the cluster
CMD /etc/init.d/ssh start && ./run.sh
\[/code\]

Now this will build and run as you would expect! Yippi! You can try this now with this command:

\[code lang=text\]
docker run --name sftestcluster -d --rm -p 19080:19080 -p 19000:19000 -p 25100-25200:25100-25200 lawrencegripper/sfonebox
\[/code\]

## Deploying into the container

So how about deploying into onto this cluster we have running? Well we could install and run the [`sfctl`](https://docs.microsoft.com/en-us/azure/service-fabric/service-fabric-application-lifecycle-sfctl) tool on our machine but keeping tools up to date and making sure they don’t clash with others is a pain so lets use docker here again.

We can create a `sfctl` docker image with a Dockerfile like so:

\[code lang=text\]
FROM python:3
RUN pip3 install sfctl
RUN sfctl cluster select --endpoint http://localhost:19080
WORKDIR /src
ENTRYPOINT \[ "bash" \]
\[/code\]

Now for some orchestration, we need to run up the container for the cluster, wait for it to become healthy, then deploy our app into is using the `sfctl` container.

We can write a nice BASH function to poll the clusters health endpoints, we use JQ to parse the output and pick out the Health and NodeCount. The function looks like this:

\[code lang=bash\]
function isClusterHealthy () {
 echo "Checking cluster status..."
 HEALTHURL="http://localhost:19080/$/GetClusterHealth?NodesHealthStateFilter=1&ApplicationsHealthStateFilter=1&EventsHealthStateFilter=1&api-version=3.0"
 HEALTH\_RESULT="$(wget --timeout=1 -qO - "$HEALTHURL" \| jq -r .AggregatedHealthState)"
 NODE\_COUNT="$(wget --timeout=1 -qO - "$HEALTHURL" \| jq -r .HealthStatistics.HealthStateCountList\[0\].HealthStateCount.OkCount)"
 echo "Current Status $HEALTH\_RESULT Nodes: $NODE\_COUNT"
 if \[ "$HEALTH\_RESULT" = "Ok" \] && \[ "$NODE\_COUNT" = "3" \]; then
 return 1
 else
 echo "Waiting for health with 3 Nodes..."
 return 0
 fi
};
\[/code\]

Now we can deploy into our cluster with a script like this:

\[code lang=bash\]
#!/bin/bash
echo "######## Upload app ###########"
sfctl application upload --path ./testapp
echo "######## Provision type ###########"
sfctl application provision --application-type-build-path testapp
echo "######## Create 200 instances ###########"
for i in {100..150}
do
 ( echo "Deploying instance $i"
 sfctl application create --app-type NodeAppType --app-version 1.0.0 --parameters "{\\"PORT\\":\\"25$i\\"}" --app-name fabric:/node25$i ) &
done
echo "Waiting for deployment to complete..."
wait
\[/code\]

We mount this into our ‘sfctl’ container and execute it like so:

\[code lang=text\]
docker run --name appinstaller -it --rm --network=host -v ${PWD}:/src lawrencegripper/sfctl -f ./uploadtestapp.sh
\[/code\]

Notice the use of `--network=host` this is crucial as without it the docker networking would prevent the `sfctl` container from connecting to the cluster container.

Note: At this point I saw a strange error from `sfct` the error was `500 code returned maximum retries exceeded` unfortunately there was no way to get the body of the 500 code through `sfctl`. Never fear [Postman](https://www.getpostman.com/) to the rescue, I manually created the REST request so I could inspect the error:

\[code lang=text\]
curl -X POST \
 'http://localhost:19080/ApplicationTypes/$/Provision?api-version=3.0&timeout=60%3FApplicationTypeImageStorePath%3Dtestapp' \
 -H 'cache-control: no-cache' \
 -H 'content-type: application/json' \
 -d '{
 "ApplicationTypeBuildPath": "testapp"

}'
\[/code\]

This returned

\[code lang=text\]
{
 "Error": {
 "Code": "FABRIC\_E\_IMAGEBUILDER\_VALIDATION\_ERROR",
 "Message": "The EntryPoint node is not found.\\nFileName: /home/ClusterDeployer/ClusterData/Data/N0010/Fabric/work/ImageBuilderProxy/AppType/NodeAppType/WebServicePkg/ServiceManifest.xml"
 }
}

\[/code\]

It hit me that node wasn’t present in the container running the cluster. I installed it with apt and then created a ‘node.sh’ file to start the app and pointed to this from my manifest. After this I got another message related to the ‘code’ folder not being found. This took a bit of time but I realized that on linux the file system is case sensitive and I had ‘Code’ not ‘code’, renaming the folder fixed that.

## Making it route requests!

Well it turns out that `./ClusterDeployer.sh` doesn’t configure things quite right for my use case. It uses the following to discover the clusters IPAddress:

\[code lang=text\]
IPAddr=\`ifconfig eth0 2>/dev/null\|awk '/inet addr:/ {print $2}'\|sed 's/addr://'\`
\[/code\]

This means our containers IPAddress is used, well with my container that isn’t routable - we need ‘localhost’ to be used.

Here is the quick fix (HACK) added to the Dockerfile

\[code lang=text\]
FROM lawrencegripper/sfonebox
RUN apt-get install nodejs -y
RUN sed -i "s%IPAddr=.\*%IPAddr=localhost%g" ClusterDeployer.sh

\[/code\]

This worked and now means the cluster publishes the addresses of our services as [http://localhost:\[port\]](http://localhost:[port]) which means Traefik can pick these up and route to them.

All done!
