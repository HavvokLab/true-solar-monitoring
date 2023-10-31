POST /solarcell-2023.*/_update_by_query
{
  "script": {
    "source": "ctx._source.owner = 'TRUE'"
  }
}