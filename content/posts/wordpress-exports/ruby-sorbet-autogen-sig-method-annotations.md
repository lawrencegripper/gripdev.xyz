---
author: gripdev
category:
  - rugby
date: "2023-01-07T08:47:09+00:00"
guid: http://blog.gripdev.xyz/?p=1582
tag:
  - rugby
  - sorbet
title: 'Ruby + Sorbet: Autogen sig method annotations'
url: /2023/01/07/ruby-sorbet-autogen-sig-method-annotations/

---
I’ve been adding Sorbet and type checking gradually to a legacy Ruby codebase dating back to the 2010‘s.

First up was getting all or most files as (a topic for another day):

```
# typed: true
```

Now I’m gradually adding method annotations, using sig annotations, to tell Sorbet what types a method accepts and returns, like so:

```
sig {params(x: SomeType, y: SomeOtherType).returns(MyReturnType)} def foo(x, y); ...; end
```

Docs: https://sorbet.org/docs/sigs

This is fairly time consuming so being lazy I wanted some help to get this done quicker.

I say “help” here as it’s never going to be perfect, with the meta programming and cruft of an old Ruby codebase it’s always going to need human validation and tweaking.

Luckily the codebase has a pretty extensive test suite so we can use that to validate the generated types match reality.

## Let’s get Autogenerating

I’m lucky to work with great engineers, George (https://hachyderm.io/@georgebrock) is one of them. He showed me a useful approach to generate sigs which saves a bunch of time.

As sorbet “knows” some return types already it can infer some sigs on methods. The trick George showed me was to get it to add those inferred sigs for you automagically.

- Set `#typed: strict` on the file, this means any methods without sigs are considered errors
- Run `srb tc --autocorrect --isolate-error-code=7017` this will tell sorbet to auto create any signatures it can work out
- Reset `#typed: true` (unless you’ve solved all errors under strict - I’m aiming for that but gradually, currently want good sigs)
- Review the auto generated sigs and make sure they’re sensible, fixing up ‘untyped’ and other issues

The cool thing here is the more sigs you have the better sorbet gets at generating the missing ones.

So what about those that can’t be created with this technique?

Well you have to write them but here there is help too. I’ve being using GitHub Copilot (https://github.com/features/copilot/) to infer or suggest sigs.

This is much more hit and miss than the first technique, it still (mostly) is a time saver but you do have to tweak the suggestions regularly.

## Playing it safe with new sigs

Now I’ve got a set of new sigs you’d think next up would be to ship them but hold fire.

Sig annotations are statically checked at dev time but also, by default, enforced at runtime too.

The danger here is the new shiny sigs you added aren’t right and your production application will start failing when they are shipped.

To work around this we need to tweak some configuration in Sorbet. In this case I configured the app to raise runtime errors when a signature wasn’t correct but only in test or in staging environment - not in production. To do this you implement a:

```
call_validation_error_handler
```

This allows you to control how sorbet reacts to a method receiving or returning a type which doesn’t match the sig annotation.

Here is the initializer I ended up with:

```
# typed: strictrequire "sorbet-runtime"# Register call_validation_error_handler callback.
# This runs every time a method with a sig fails to type check at runtime.
# See: https://sorbet.org/docs/runtime#on_failure-changing-what-happens-on-runtime-errors# In all environments report a sig violation to Sentry.
# In any non-production environment raise an error if a sig is violated.
T::Configuration.call_validation_error_handler = lambda do |signature, opts|
  failure_message = opts[:pretty_message]
  Scrolls.log(at: :sorbet_runtime_sig_checking, msg: failure_message.squish)
  error = TypeError.new(failure_message)
  track_error_in_sentry!(error)
  raise error unless Rails.env.production?
end

```

Docs: https://sorbet.org/docs/runtime#on\_failure-changing-what-happens-on-runtime-errors

There you go, ship the new code with its sigs and keep an eye on Sentry to see if any need tweaking based on real production usage without those causing production errors.
