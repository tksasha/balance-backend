# frozen_string_literal: true

RSpec.describe 'GetItem' do
  describe 'GET /items/{id}.json' do
    context 'when item is exist' do
      let(:response) { connection.get("/items/#{ item.id }") }

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

      it { expect(status).to eq(:ok) }

      it { expect(body).to eq(expected) }
    end

    context 'when item is not found' do
      let(:response) { connection.get('/items/0') }

      it { expect(status).to eq(:not_found) }

      it { expect(body).to eq errors: { base: ['resource not found'] } }
    end

    context 'when id is a string' do
      let(:response) { connection.get('/items/string') }

      it_behaves_like 'show.json'

      it { expect(status).to eq(:bad_request) }

      it { expect(body).to eq errors: { id: ['is invalid'] } }
    end
  end
end
