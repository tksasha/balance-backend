# frozen_string_literal: true

require 'active_record'
require 'sqlite3'

ActiveRecord::Base.establish_connection(adapter: 'sqlite3', database: '../db/test.sqlite3')

class Category < ActiveRecord::Base; end

class Item < ActiveRecord::Base
  belongs_to :category
end

RSpec.configure do |config|
  config.after(:suite) do
    Item.delete_all
    Category.delete_all
  end
end
