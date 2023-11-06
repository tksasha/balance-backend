# frozen_string_literal: true

RSpec.describe 'GetItemsList' do
  subject { connection.get('/items') }

  its(:status) { is_expected.to eq 200 }

  its('body.chomp') { is_expected.to eq '"Items List"' }
end
