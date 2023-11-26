# frozen_string_literal: true

RSpec.describe 'GetItem' do
  before do
    db.execute('DELETE FROM items')

    sql = <<-SQL
      INSERT INTO items(id, date, formula, sum, category_id, description)
      VALUES(1512, "2023-11-26", "42.1 + 69.01", 111.11, 2919, "lorem ipsum ...")
    SQL

    db.execute(sql)
  end

  after { db.execute('DELETE FROM items') }

  context 'when everything is fine' do
    subject { connection.get('/items/1512') }

    let(:expected) do
      {
        id: 1512,
        date: '2023-11-26',
        sum: 111.11,
        category_id: 2919,
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
