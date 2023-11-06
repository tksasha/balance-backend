# frozen_string_literal: true

require 'faraday'
require 'rspec/its'

RSpec.configure do |config|
  config.order = :random
end

def connection
  Faraday.new(url: 'http://localhost:3000') do |builder|
    builder.request :json
    builder.response :json
  end
end
