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

def body
  response&.body.presence&.deep_symbolize_keys || {}
end

def status
  {
    200 => :ok,
    201 => :created,
    400 => :bad_request,
    404 => :not_found
  }[response.status]
end
