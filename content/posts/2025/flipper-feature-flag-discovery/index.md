---
title: "Ruby: Flipper Feature Flag Discovery with AST parsing"
date: 2025-01-31T07:02:00+00:00
author: "Lawrence Gripper"
tags: ["programming", "ruby", "ast", "feature-flags", "flipper"]
categories: ["programming"]
url: /2025/01/31/flipper-feature-flag-discovery/
draft: false
---

So you have Flipper in your ruby codebase for handling feature flags but... now you have used them and need
a way to find them and help you pay down the tech debt of cleaning them up.

Well good news, you can use AST Parsing to find all the instance.

This simple parser walks all the Ruby code looking for calls to `Flipper.enabled?` and makes a list of the features flags used.

```ruby
    class FlipperEnableFinder < Parser::AST::Processor
      include RuboCop::AST::Traversal

      sig { returns(T::Array[Symbol]) }
      attr_reader :found_features

      PATTERN = <<~PATTERN
      (send (const {nil? (cbase)} :Flipper) :enabled? sym ...)
    PATTERN

      sig { void }
      def initialize
        @found_features = T.let([], T::Array[Symbol])
      end


      sig { params(node: T.untyped).returns(T.untyped) }
      def on_send(node)
        if RuboCop::AST::NodePattern.new(PATTERN).match(node)
          feature = node.children.detect { |c| c.respond_to?(:type) && c.type == :sym }.children.first
          @found_features << feature
        end
      end
    end
```

So then all we need to do is pull in all the ruby code and pass it through the AST rule.

This is a little slow so using `parallel_ruby` helps us out here and some pre-processing to get a smaller list of Ruby files to parse with our AST processor.

Full implementation looks like this:

```ruby
class Usage
    # Skip dirs that don't have any ruby
    EXCLUDED_DIRS = %w[.npm coverage .git bin public tmp vendor node_modules script spec sorbet]

    class FlipperEnableFinder < Parser::AST::Processor
      include RuboCop::AST::Traversal

      sig { returns(T::Array[Symbol]) }
      attr_reader :found_features

      PATTERN = <<~PATTERN
      (send (const {nil? (cbase)} :Flipper) :enabled? sym ...)
    PATTERN

      sig { void }
      def initialize
        @found_features = T.let([], T::Array[Symbol])
      end


      sig { params(node: T.untyped).returns(T.untyped) }
      def on_send(node)
        if RuboCop::AST::NodePattern.new(PATTERN).match(node)
          feature = node.children.detect { |c| c.respond_to?(:type) && c.type == :sym }.children.first
          @found_features << feature
        end
      end
    end

    # Cache the result once you have it as its expensive to compute
    @@flipper_features_result = T.let(nil, T.nilable(T::Array[Symbol]))

    sig { returns(T::Array[Symbol]) }
    def self.find_all_from_code
      return @@flipper_features_result unless @@flipper_features_result.nil?
      
      # As AST parsing is expensive, do a first pass with glob and checking content as these
      # are cheap
      files_with_flags_used = Dir.glob("./**/*.rb").filter_map { |f|
        dirs = Pathname.new(f).dirname.to_s.split("/")
        next if dirs.any? { |dir| EXCLUDED_DIRS.include?(dir) }
        content = File.read(f)
        content if content.match(/Flipper.enabled\?.*/)
      }

      # Now we have a small list lets parse the AST and find our flags
      found_features = Parallel.map(files_with_flags_used, in_processes: 32) do |file|
        rule = FlipperEnableFinder.new
        source = RuboCop::AST::ProcessedSource.new(file, 3.1)
        source.ast.each_node do |node|
          rule.process(node)
        end
        rule.found_features
      end

      @@flipper_features_result = found_features.flatten.uniq

      return @@flipper_features_result
    end
  end
```