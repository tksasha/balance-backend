# frozen_string_literal: true

RSpec.describe 'UpdateItem' do
  subject { connection.patch('/items/42') }

  it { expect(status).to eq 200 }

  it { expect(content_type).to eq 'application/json' }

  it { expect(body).to eq 'Item Updated' }
end
