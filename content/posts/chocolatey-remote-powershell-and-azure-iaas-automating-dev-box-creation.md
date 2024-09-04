---
author: gripdev
category:
  - how-to
date: "2013-07-07T11:33:42+00:00"
guid: http://gripdev.wordpress.com/?p=340
title: Chocolatey, Remote PowerShell and Azure IAAS - automating dev box creation
url: /2013/07/07/chocolatey-remote-powershell-and-azure-iaas-automating-dev-box-creation/

---
Hi All,

Recently I've needed a throw away environment and so turned to Azure IAAS to get a nice VM up and running quickly. To do this I've turned to Chocolatey, its a great utility for installing application and I'd been using it to install simple bits and bobs on my local machine. I fell in love, there are some caveats to be aware of but in short its awesome.

So I started playing, the aim - Automate the create and installation of an Azure VM with all the dev tools I needed with no user interaction.

1. I came across a great post by [Michael Washam](http://michaelwasham.com/windows-azure-powershell-reference-guide/introduction-remote-powershell-with-windows-azure/ "Michael Washam") which details how to boot a Azure VM then initiate remote powershell onto the box. This is awesome, it let me spin up a box and set it off installing the windows features.
1. I tweaked this script a bit to remove the need for the credentials prompt and added a quick check to ensure I had the right privileges on the box for it to run correctly.
1. Once the Windows Features are finished I install Chocolatey and add in all the nice programs that you'd like on your dev box. Vs2012, Linqpad, notepad++ etc
1. Chocolatey also lets you install from the webplatform installer so I then set about installing all the stuff I need from webpi. AzurePowershell, SQLExpressTools et
1. Wait for the script to complete and all being well you've got a nice devbox ready to go!
1. For good measure the script kicks off a RDP session to the box and interactive remote powershell to finish.

\[code language="Powershell"\]

If (-NOT (\[Security.Principal.WindowsPrincipal\] \[Security.Principal.WindowsIdentity\]::GetCurrent()).IsInRole(\`
 \[Security.Principal.WindowsBuiltInRole\] "Administrator"))
{
 Write-Warning "You do not have Administrator rights to run this script!\`nPlease re-run this script as an Administrator as it will need to install the remote certificate."
 Break
}

$ErrorActionPreference = "Stop"
$user = "usernamegoeshere"
$pwd = "passwordgoeshere"
$SecurePassword = $pwd \| ConvertTo-SecureString -AsPlainText -Force
$svcName = "servicenamegoeshere"
$VMName = "vmnamegoeshere"
$location = "North Europe" #change as needed
$imageName = "a699494373c04fc0bc8f2bb1389d6106\_\_Windows-Server-2012-R2-Preview-201306.01-en.us-127GB.vhd" #or whichever azure image you'd like as the base.
$credential = new-object -typename System.Management.Automation.PSCredential -argumentlist $user,$SecurePassword

New-AzureVMConfig -Name $VMName -InstanceSize "ExtraSmall" -ImageName $imageName \|
 Add-AzureProvisioningConfig -Windows -AdminUsername $user -Password $pwd \|
 Add-AzureEndpoint -Name "http" -Protocol tcp -LocalPort 80 -PublicPort 80 \|
 New-AzureVM -ServiceName $svcName -Location $location -WaitForBoot
function InstallWinRMCert($serviceName, $vmname)
{
 $winRMCert = (Get-AzureVM -ServiceName $serviceName -Name $vmname \| select -ExpandProperty vm).DefaultWinRMCertificateThumbprint

 $AzureX509cert = Get-AzureCertificate -ServiceName $serviceName -Thumbprint $winRMCert -ThumbprintAlgorithm sha1

 $certTempFile = \[IO.Path\]::GetTempFileName()
 $AzureX509cert.Data \| Out-File $certTempFile

 # Target The Cert That Needs To Be Imported
 $CertToImport = New-Object System.Security.Cryptography.X509Certificates.X509Certificate2 $certTempFile

 $store = New-Object System.Security.Cryptography.X509Certificates.X509Store "Root", "LocalMachine"
 $store.Certificates.Count
 $store.Open(\[System.Security.Cryptography.X509Certificates.OpenFlags\]::ReadWrite)
 $store.Add($CertToImport)
 $store.Close()

 Remove-Item $certTempFile
}

\# Get the RemotePS/WinRM Uri to connect to
$uri = Get-AzureWinRMUri -ServiceName $svcName -Name $VMName

\# Using generated certs – use helper function to download and install generated cert.
InstallWinRMCert $svcName $VMName

\# Use native PowerShell Cmdlet to execute a script block on the remote virtual machine
Invoke-Command -ConnectionUri $uri.ToString() -Credential $credential -ScriptBlock {
 write-host "Installing windows features"
 $logLabel = $((get-date).ToString("yyyyMMddHHmmss"))
 $logPath = "$env:TEMPinit-webservervm\_webserver\_install\_log\_$logLabel.txt"
 Import-Module -Name ServerManager

 #example features
 Install-WindowsFeature -Name Web-Server -IncludeManagementTools -LogPath $logPath

 #nb This feature is required for webpi installer to function correctly
 install-windowsfeature -name NET-Framework-Features -LogPath $logPath

write-host "Installling Choc"

iex ((new-object net.webclient).DownloadString("http://chocolatey.org/install.ps1"))

write-host "installing dev tooling"
 #example packages
 cinst VisualStudio2012Ultimate
 cinst MsSqlServer2012Express
 cinst linqpad4
 cinst notepad++
 cinst SQLManagementStudio -source webpi
 cinst SQLExpressTools -source webpi
 cinst VWDOrVs11AzurePack -source webpi
 cinst WindowsAzurePowershell -source webpi
}
Write-host "Install command finished - Launching remote desktop and interactive remote powershell session to devbox"
Get-AzureRemoteDesktopFile -ServiceName $svcName -Name $VMName –Launch
Enter-PSSession -ConnectionUri $uri.ToString() -Credential $credential
\[/code\]
