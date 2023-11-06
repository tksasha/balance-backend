# frozen_string_literal: true

RSpec.describe 'CreateItem' do
  subject { connection.post('/items') }

  it { expect(status).to eq 201 }
end
