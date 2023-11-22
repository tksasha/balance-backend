# frozen_string_literal: true

require 'active_support/core_ext/hash/keys'
require 'active_support/core_ext/hash/slice'
require 'debug'
require 'faraday'
require 'sqlite3'

RSpec.configure do |config|
  config.order = :random
end

def connection
  url = 'http://localhost:3000'

  headers = { 'Content-Type': 'application/json' }

  Faraday.new(url:, headers:) do |builder|
    builder.request :json
    builder.response :json
  end
end

def status
  subject.status
end

def content_type
  subject.headers[:content_type]
end

def body
  subject.body.symbolize_keys
end

alias responsed body

def db
  @db ||= SQLite3::Database.new('../db/test.sqlite3')
end
