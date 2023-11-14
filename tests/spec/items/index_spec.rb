# frozen_string_literal: true

RSpec.describe 'GetItemsList' do
  subject { connection.get('/items') }

  it { expect(status).to eq 200 }

  it { expect(content_type).to eq 'application/json' }

  it { expect(body.size).to eq 10 }
end
