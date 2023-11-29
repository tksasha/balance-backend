# frozen_string_literal: true

RSpec.shared_examples 'show.json' do
  let(:body) { response.body.deep_symbolize_keys }

  it { expect(response[:content_type]).to eq 'application/vnd.api+json' }
end
