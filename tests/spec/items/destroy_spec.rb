# frozen_string_literal: true

RSpec.describe 'DeleteItem' do
  subject { connection.delete('/items/1') }

  it { expect(status).to eq 200 }

  xit { expect(content_type).to eq 'application/json' }

  xit { expect(body).to eq 'Item Deleted' }
end
