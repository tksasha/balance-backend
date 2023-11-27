# frozen_string_literal: true

RSpec.describe 'GetItem' do
  context 'when everything is fine' do
    subject { connection.get("/items/#{ item.id }") }

    let!(:item) { create(:item) }

    let(:expected) do
      {
        id: item.id,
        date: '2023-11-26',
        sum: 111.11,
        category_id: item.category_id,
        formula: '42.1 + 69.01',
        description: 'lorem ipsum ...'
      }
    end

    it { expect(status).to eq 200 }

    it { expect(content_type).to eq 'application/json' }

    it { expect(responsed).to eq expected }
  end

  context 'when item is not found' do
    subject { connection.get('/items/0') }

    it { expect(status).to eq 404 }

    it { expect(content_type).to eq 'application/json' }
  end

  context 'when id is a string' do
    subject { connection.get('/items/string') }

    it { expect(status).to eq 404 }

    it { expect(content_type).to eq 'application/json' }
  end
end
