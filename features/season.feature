Feature: Crops by season

    Scenario: All crops are listed grouped by season when no season is provided
        When season is requested without a specified season
        Then all crops grouped by season must be returned
        
    Scenario: A specified season only returns data for that season
        When the summer season's crops is requested
        Then only the summer crops must be returned

    Scenario: An unmatched season returns an informative message
        When the breakfast season's crops are requested
        Then a message informing no matched season was found must be returned
