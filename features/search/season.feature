Feature: Search by Season

  Scenario: A season specified returns the crops that grow in that season
  # stardew-crops search --season fall
  When a search by season for fall is performed
  Then a list of crops that grow in fall must be returned

  Scenario: An unmatched season returns an applicable error message
  # stardew-crops search --season banana
  When a search by season for banana is performed
  Then an error indicating no matches found must be returned

  Scenario: Searching by season without specifying one returns an error message
  # stardew-crops search --season
  When a search by season is performed with no value
  Then an error indicating a value must be specified must be returned
