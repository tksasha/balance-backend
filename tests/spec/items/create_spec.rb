# frozen_string_literal: true

RSpec.describe 'CreateItem' do
  subject { connection.post('/items') }

  it { expect(status).to eq 201 }

  it { expect(content_type).to eq 'application/json' }

  it { expect(body).to eq 'Item Created' }
end
