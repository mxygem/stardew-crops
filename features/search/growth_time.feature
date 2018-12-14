Feature: Search by Growth Time

    * Search flags of growthlt and growthgt are available for specifying times
    * Time is calculated as less/greater than or equal to

    Scenario: A less than time of 5 is specified
        # stardew-crops search --growthlt 5
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
