# frozen_string_literal: true

RSpec.describe 'DeleteItem' do
  subject { connection.delete('/items/1') }

  it { expect(status).to eq 200 }

  it { expect(content_type).to eq 'application/json' }

  it { expect(body).to eq 'Item Deleted' }
end
