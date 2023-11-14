# frozen_string_literal: true

RSpec.describe 'GetItem' do
  context 'when everything is fine' do
    subject { connection.get('/items/42') }

    it { expect(status).to eq 200 }

    it { expect(content_type).to eq 'application/json' }

    xit { expect(body).to eq('id' => 42, 'name' => 'Pretty Red Dress') }
  end

  context 'when item is not found' do
    subject { connection.get('/items/0') }

    it { expect(status).to eq 404 }

    it { expect(content_type).to eq 'application/json' }

    it { expect(body).to eq 'Not Found' }
  end

  context 'when id is invalid' do
    subject { connection.get('/items/invalid') }

    it { expect(status).to eq 422 }

    it { expect(content_type).to eq 'application/json' }

    it { expect(body).to eq 'Unprocessable Entity' }
  end
end
