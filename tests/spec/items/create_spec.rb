# frozen_string_literal: true

RSpec.describe 'CreateItem' do
  subject { connection.post('/items', requested) }

  let(:requested) do
    {
      date: '2023-11-18',
      category_id: 2919,
      formula: '42.1 + 69.01',
      description: 'Lorem ipsum dolor sit amet'
    }
  end

  let(:expected) do
    {
      id: 1,
      date: '2023-11-18',
      sum: 111.11,
      category_id: 2919,
      formula: '42.1 + 69.01',
      description: 'Lorem ipsum dolor sit amet'
    }
  end

  before { db.execute('DELETE FROM items') }

  context 'when params are valid' do
    it { expect(status).to eq 201 }

    it { expect(content_type).to eq 'application/json' }

    it { expect(responsed).to eq expected }
  end

  context 'when id is specified' do
    let(:requested) do
      {
        id: 1211,
        date: '2023-11-22',
        category_id: 2919,
        formula: '42.1 + 69.01',
        description: 'Lorem ipsum dolor sit amet'
      }
    end

    let(:expected) do
      {
        id: 1,
        date: '2023-11-22',
        sum: 111.11,
        category_id: 2919,
        formula: '42.1 + 69.01',
        description: 'Lorem ipsum dolor sit amet'
      }
    end

    it { expect(responsed).to eq expected }
  end
end
