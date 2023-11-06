# frozen_string_literal: true

RSpec.describe 'UpdateItem' do
  subject { connection.patch('/items/42') }

  it { expect(status).to eq 200 }

  xit { expect(body).to eq 'Update Item' }
end
