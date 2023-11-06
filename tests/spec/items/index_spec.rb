# frozen_string_literal: true

RSpec.describe 'GetItemsList' do
  subject { connection.get('/items') }

  it { expect(status).to eq 200 }

  xit { expect(content_type).to eq 'application/json' }

  xit { expect(body).to eq '"Items List"' }
end
