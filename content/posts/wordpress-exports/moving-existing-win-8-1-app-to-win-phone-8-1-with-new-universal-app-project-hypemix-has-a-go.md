---
author: gripdev
category:
  - windows-store
date: "2014-04-07T01:00:02+00:00"
guid: http://gripdev.wordpress.com/?p=519
title: Moving existing Win 8.1 App to Win Phone 8.1 with new Universal App project – HypeMix has a go!
url: /2014/04/07/moving-existing-win-8-1-app-to-win-phone-8-1-with-new-universal-app-project-hypemix-has-a-go/

---
This a quick and dirty write up of how I moved my existing Windows 8.1 App to Windows Phone 8.1 in a weekend. Now before we get into that I'm going to be honest and say there is still some work to do around the UI, as playing music on the phone is slightly different to a PC or Tablet. What I'm talking about here is all the rest of the work.

So, I set out on my ambition goal with a simple step. My current solution had 4 Projects, one for the app, one for the API that it uses, one for tests and a second for tests just to be safe. If you want to play with the app you can get it here - http://apps.microsoft.com/windows/en-gb/app/hypemix/f5e81d5d-b972-456c-9f33-655c23488815

## Step 1:

Right click on the Win 8.1 Project and Click "Add windows Phone project"

This gives you a "Shared" section, if you've worked with linked files before this is along those lines. When you compile each of the Win or WinPhone projects, the files in shared are included.

## Step 2:

![](/wp-content/uploads/2014/04/040714_0825_movingtouni1.png)

Now that you've got this lovely shared section you can move over all the obvious bits that you want to share. To kick this off I moved over my Commands, Exceptions, Services and ViewModels.

(To clarify the app uses Reactive Extensions to create a Publish and Subscribe architecture. When you click a button this binds to a Command on a ViewModel. The Command publishes an Event and these events are subscribed to by one or more services.

For example, Clicking "Play" invokes the "PlayPauseCommand" which in turn publishes a "PlayPause" Event. This event is subscribed to by the "MediaService", when received the service starts playing the track. It's also subscribed to by the "NotificationService" which pops up a toast to inform the user about which track is playing)

The first awesome thing I saw here was that my existing Win8.1 app still worked fine, moving the stuff to the shared section hadn't had any effect on it. Big win, no changes needed to existing app so far!

The second awesome thing I saw here was EVERY BIT OF CODE, not counting the XAML and code behind ~8 files, was shared. That's right, EVERY ONE OF THE 1.2K LINES OF CODE in the Windows 8.1 Solution could be shared. So far so good, let's see how the API project gets on.

## Step 3:

I mentioned earlier that I have a Windows 8.1 Class Library, which implements the HypeM Api. Now in the new SDK we have Universal App - Portable Class library. This is just a PCL which targets Win8.1 and WinPhone 8.1. So I can take advantage of this I needed to move all my stuff over to the new project type.

[@saxenanavit](https://twitter.com/saxenanavit) let me know that you can simply retarget your project to Windows Phone 8.1 by going into properties and selecting Windows Phone. As this is a much simpler than using linked files I've cross out my original bit and moved to using this!

[![Bko0vYICQAAo-aD](/wp-content/uploads/2014/04/bko0vyicqaao-ad.png?w=300)](/wp-content/uploads/2014/04/bko0vyicqaao-ad.png)To do this I created a new "Universal App Portable Class Library".![](/wp-content/uploads/2014/04/040714_0825_movingtouni2.png)Then I used project linker to link all the files and add them into the new project type. (I had to download the VSIX file and alter the manifest so that it would work with VS2013 but once I'd done that it was all good, [VSIX below](https://onedrive.live.com/redir?resid=D11EE8A531F0B903!61939&authkey=!APXZCliKOjuoeV4&ithint=file%2c.vsix))

After re-targeting the library I hit my first snag. When I happily hit F6, expecting the same level of success, my heart sunk as I saw errors. "Here we go" I thought, I'm going to have to refactor loads of code.

Once I'd talked myself back from the edge and actually looked at the errors I found there was only one culprit.

This was the credentials picker, which isn't currently in the shared WinRT Core.

using Windows.Security.Credentials.UI;
`CredentialPickerOptions opts = new CredentialPickerOptions(); `This is the awesome bit, apart from that everything else just worked! That's another 1.4k+ lines of shared code.

Once again I jumped back to the Windows 8.1 app and pressed F5 and was happy to see the existing app still working well, with its new projects and shared core.

# Next – Build the Windows Phone 8.1 Project

I switched over and tried to build the Windows Phone 8.1 Project, the errors quickly stacked up and I had another "Oh this is going to be tough" moment.

Actually it wasn't, in step 2 I moved some stuff to the shared project. The odd one here is that any references required for this code need to be referenced in both the Windows Phone and Windows 8.1 project. This is because, as I mentioned, the Shared bit isn't actually a project, it's a bunch of files and as such it doesn't have references.

Now normally I'd just jump on nuget to get hold of these references, which is where the second bit of trouble struck …

![](/wp-content/uploads/2014/04/040714_0825_movingtouni4.png)

I'd used the Prism for Store Apps package for a number of things in the Win8.1 app. When I tried to add it from Nuget I got some nasty red telling me it wasn't supported by this project type.

I thought to myself, now WinRT is basically shared I wonder if I can use the Store version. Sure enough add it manually and you're golden!

![](/wp-content/uploads/2014/04/040714_0825_movingtouni5.png)

The second roadblock I hit was around App Insights, I couldn't get this to play nice with Phone 8.1. However as I'm using a nice Pub/Sub model for events I can just leave this be for the time being while they're SDK gets updated.

I will look to switch out the "AppInsightService" below with a WinPhone friendly one once it becomes available.

`public class AppInsightService``   { ``       public AppInsightService() ``       { ``           EventPubSub.Subscribe<CurrentPlaylistUpdate>(x => ``           { ``               var properties = new Dictionary<string, object>() ``{ { "artist", x.CurrentPlayingTrack.artist }, { "track", x.CurrentPlayingTrack.title }, { "id", x.CurrentPlayingTrack.itemid } }; ``               ClientAnalyticsChannel.Default.LogEvent("Updated Song", properties); ``           }); `

So I started a little platform specific folder in my Win 8 Project. Now containing app insights and ApiAuth, which uses the Credentials UI mentioned earlier.

![](/wp-content/uploads/2014/04/040714_0825_movingtouni6.png)

Then I had some fun sorting out some more references from the Windows Phone Side, I used Unity for my decency dependency injection and needed to get this up and running in the Phone Project. Again I saw issues with Nuget not playing nice with the new Project type.

![](/wp-content/uploads/2014/04/040714_0825_movingtouni7.png)

So I jumped into my "packages" file and found the NetCore45 version that was currently used in the main Windows 8.1 App.

![](/wp-content/uploads/2014/04/040714_0825_movingtouni8.png)

Adding this manually and again I was all good!

# Unity and Next set of errors….

 [![Errors1](/wp-content/uploads/2014/04/errors1.png?w=300)](/wp-content/uploads/2014/04/errors1.png)

In my Windows 8.1 App I have a static reference to the Unity container in the App.xaml.cs. Rightly or wrongly I configure the container in the App.Xaml.Cs and use this to keep a static reference to it while the app is running.

![](/wp-content/uploads/2014/04/040714_0825_movingtouni10.png)

I could see my Windows 8.1 Project building fine but not my Windows Phone project. This was because, as the Win8.1 builds, it pulls in the shared files from the new Shared area and all the references to "App.Container" are present in its App.Xaml.Cs.

When I build the Windows Phone 8.1 the errors appear as the Phone project doesn't have that container setup in its app.xaml.cs. This was a good lesson, be careful when using static stuff which hangs off the App object in your shared section, as you'll have different versions in each project!

Actually, in this case, it's a nice thing as it makes perfect sense to have different Unity config for each app. For example, one of the items that unity resolves in my ApiAuth.cs - which we've had to move into platform specific folder in the Windows 8 Project as it couldn't be shared. So in each App.Xaml.Cs I can have different unity config which provides a Platform specific version of that class.

# Navigation

Next tricksy one is navigation, as you can see from these error messages I'm using a static reference to the rootframe to move around. This isn't going to do the trick, as we're going to have different types depending on which version we're in because the Pages don't sit in the Shared Section.

[![Errors](/wp-content/uploads/2014/04/errors.png?w=300)](/wp-content/uploads/2014/04/errors.png)

My first thought here was to use Unity, as I've already got the way paved for me. I was going to implement a quick Interface and use the unity configuration to inject a WinPhone of Win8.1 Navigation handler in each of the projects. That way I drop this nasty stuff and I can map a navigate event to a different page based on the platform etc.

When I thought about it a bit more though I decided to go with leveraging the existing publish subscribe model, with services, I'm using in the app.

I created a "NavigationEvent" with "param" and "PageName" then a service which subscribes to these events and kicks off the navigation events.

`namespace HyperM8.Shared.PubSub ``{ ``    public class NavigationEvent``    { ``        public string PageName { get; set; } ``        public string ParamsString { get; set; } ``    } ``} `This led to changing around 5 or 6 of my commands to using this model.

`var navEvent = new NavigationEvent(){ PageName = "MoreItemsPage", ParamsString = string.Format("{0},{1}", url, blog.sitename)}; ``   EventPubSub.Publish<NavigationEvent>(navEvent); `

VS the existing direct reference to the rootFrame to move around.

```
App.RootFrame.Navigate(typeof(MoreItemsPage), string.Format("{0},{1}",song.PlaylistUrl, song.SongCategory));

```

(What I also realised, during this process, was that I should be injecting my pubsub bus "EventPubSub" as an Interface so I can mock it out and test with something like MOQ. So a side effect of some of this refactoring was that I ended up breaking some of my tests.)

Next step was to implement the navigation service in the platform specific folder on the Windows 8.1 project.

I'm using a nice bit of functionality in Unity here called AllClasses. This allowed me to interrogate the classes loaded in the application and find a type that matched the "PageName" passed into the "NavigationEvent" then navigate to that page.

`public NavigationService() ``{ ``    EventPubSub.Subscribe<NavigationEvent>(navEvent => ``    { ``        var type = AllClasses.FromApplication().Where(c => c.Name == navEvent.PageName).FirstOrDefault(); ``        if (type == null) ``        { ``            Debug.WriteLine("Failed to navigate to {0}, type not found", navEvent.PageName); ``            return; ``        } ``        App.RootFrame.Navigate(type, navEvent.ParamsString); ``    }); ``} `(@robgarfoot has pointed out that this is quite a bit of overhead, going through every class in ever assembly in the app each time a nav event is fired. So far I haven't seen a huge hit from this but I do agree this is a bit smelly. When I have a bit more time I'm going to look at using Unity to make this a once time mapping, rather than something that happens every time the app navigates)

So with this now in place and all the commands updated, my project builds and my Windows 8.1 app is still working nicely. What I then realised was that I could put this "NavigationService" into my Shared section and as long as I named the pages correctly it would all work and also be shared.

# Summary

I've shared a HUGE AMOUNT OF CODE and I'm way ahead where I would have been if I'd tried this before this updated. The only bits where I saw issues was with Navigation, Nuget Packages, References and Credentials Picks… none of which took me more than 15mins each to get around…

In short…. THIS IS AWESOME!!

My project structure is as follows, unloaded have either been killed by my changes or legacy.

![](/wp-content/uploads/2014/04/040714_0825_movingtouni12.png)

Next up was migrating some of the XAML for the HubPage. Certainly there are more differences here. I got a massive helping hand from the new approach to Universal apps and project. I took a chance and moved my Resource Dictionary, containing styles and Data Templates, into the Shared Project.

![](/wp-content/uploads/2014/04/040714_0825_movingtouni13.png)

Next I updated the app.xaml in both Win8.1 and WinPhone to pick up these styles and merge them into the dictionary. (Yes, there is a file called MupetStyles. No, I don't have a clue why I called it that either).

`<Application x:Class="HypeM8.App" xmlns="http://schemas.microsoft.com/winfx/2006/xaml/presentation" xmlns:x="http://schemas.microsoft.com/winfx/2006/xaml" xmlns:local="using:HypeM8">``    <Application.Resources>``        <ResourceDictionary>``            <ResourceDictionary.MergedDictionaries>``                <ResourceDictionary Source="StandardStyles.xaml" />``                <ResourceDictionary Source="MupetStyles.xaml" />``            </ResourceDictionary.MergedDictionaries>``        </ResourceDictionary>``    </Application.Resources>``</Application>`

To my absolute delight, this worked and my styles came across. Emboldened by this I picked up the body of my existing Windows 8.1 HubPage.xaml and Ctrl-C, Ctrl-V'd it into a page on the phone. With the exception of about 15 errors relating to slight differences in the XAML, THIS ALSO WORKED! I EVEN GOT MY NICE DESIGN TIME EXPEREINCE!

Just to emphasize this, I'd shared all my converters, styles and most of the page XAML between the projects. I still am completely blown away by this!

![](/wp-content/uploads/2014/04/040714_0825_movingtouni14.png)

I was now at the stage where my app was building, had a working design time view and also nearly everything was shared (API, Converters, Services, Styles, ViewModels…)

# Next – F5 onto the Emulator and see if it works for real

So by this point, if you're anything like me, you're thinking it's all going to go wrong. And it did for a little bit, I came up against this cryptic little error while working with a message dialog.

![](/wp-content/uploads/2014/04/040714_0825_movingtouni15.png)

After about 20mins of trying out different stuff I worked out that this actually means "MessageDialog" with more than two buttons isn't supported on the phone 8.1. With that all ironed out I eagerly awaiting the next F5.

![](/wp-content/uploads/2014/04/040714_0825_movingtouni16.png)

There I was, my app ported over with the first page up and running. Sure I needed to do some work on tweaking styling, porting other xaml pages and implementing a login screen but apart from that I was there. I'd shared nearly 2k lines of code and got to this point inside 2 days.

Hopefully this was a good guide for others trying to do the same, keep an eye out for HypeMix in the Windows Phone 8.1 Store once I've tweaked the styling and finished the rest of the pages. Depending on how I get on I may do a follow up post.
