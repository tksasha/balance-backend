# frozen_string_literal: true

RSpec.describe 'GetItem' do
  context 'when everything is fine' do
    subject { connection.get('/items/42') }

    it { expect(subject.status).to eq 200 }

    it { expect(subject.body).to eq('id' => 42, 'name' => 'Pretty Red Dress') }

    it { expect(subject.headers[:content_type]).to eq 'application/json' }
  end

  context 'when item is not found' do
    subject { connection.get('/items/0') }

    its(:status) { is_expected.to eq 404 }

    its('body.chomp') { is_expected.to eq '"Not Found"' }
  end
end
