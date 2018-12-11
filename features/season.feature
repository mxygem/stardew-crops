Feature: Crops by season

    Scenario: All crops are listed grouped by season when no season is provided
        When season is requested
        Then the output must match `{"spring":["garlic","potato"],"summer":["blueberry","radish","starfruit"],"fall":["cranberries","yams"]}`
     
    Scenario: A specified season only returns data for that season
        When the summer season is requested
        Then the output must match `{"summer":["blueberry","radish","starfruit"]}`

    Scenario: An unmatched season returns a message informing the user of such
        When the breakfast season is requested
        Then the output must match "Unknown season for breakfast"
