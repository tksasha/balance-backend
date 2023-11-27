# frozen_string_literal: true

FactoryBot.define do
  factory :category do
    name { Faker::Commerce.department(max: 6) }
  end

  factory :item do
    date { '2023-11-26' }
    formula { '42.1 + 69.01' }
    sum { 111.11 }
    category
    description { 'lorem ipsum ...' }
  end
end
