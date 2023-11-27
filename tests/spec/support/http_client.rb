# frozen_string_literal: true

require 'debug'
require 'faraday'

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
  subject.body
end

def responsed
  body.symbolize_keys
end
