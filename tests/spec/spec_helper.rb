# frozen_string_literal: true

require 'factory_bot'
require 'faker'

Dir[File.join('.', 'spec', 'support', '**', '*.rb')].each { |f| require f }

RSpec.configure do |config|
  config.order = :random

  config.include FactoryBot::Syntax::Methods

  config.before(:suite) do
    FactoryBot.find_definitions
  end
end
