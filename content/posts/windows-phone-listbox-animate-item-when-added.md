---
author: gripdev
category:
  - how-to
date: "2013-06-22T13:49:25+00:00"
guid: http://gripdev.wordpress.com/?p=264
title: Windows Phone - Listbox animate item when added
url: /2013/06/22/windows-phone-listbox-animate-item-when-added/

---
Really quick one that might help out others. This datatemplate is used in my listbox as follows:

\[code language="xml"\]
<ListBox ItemsSource="{Binding example}" ItemTemplate="{StaticResource ListNewsItem}"  ... />
\[/code\]

then the template looks as follows, mine is wrapped in a button as I then bind commands but I've removed that for this example.

\[code language="xml"\]<DataTemplate x:Key="ListNewsItem">
<Button x:Name="newsItemBtn">
<Button.Resources>
<EventTrigger x:Name="event" RoutedEvent="Canvas.Loaded">
<BeginStoryboard>
<Storyboard x:Name="FadeIn">
<DoubleAnimationUsingKeyFrames Storyboard.TargetProperty="(UIElement.Opacity)" Storyboard.TargetName="newsItemBtn">
<EasingDoubleKeyFrame KeyTime="0" Value="0.01"/>
<EasingDoubleKeyFrame KeyTime="0:0:0.5" Value="1"/>
</DoubleAnimationUsingKeyFrames>
</Storyboard>
</BeginStoryboard>
</EventTrigger>
</Button.Resources>
<Button.Style>
<StaticResource ResourceKey="InvisibleButtonStyle"/>
</Button.Style>

<TextBlock Text="{Binding ItemBindingExample}"/>
</Button>
</DataTemplate>
\[/code\]

You now get a nice fade in when the an item is added to the list rather than a rough jump, looks much nicer!
