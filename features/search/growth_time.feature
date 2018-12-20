Feature: Search by Growth Time

    * Search flags of growth, growthlt, and growthgt
        are available for specifying times
    * Time for the lt and gt flags are calculated as
        less/greater than or equal to with time for
        the growth flag being exact

    Scenario: A time of 13 is specified
        # stardew-crops search --growth 13
        When a search by growth for 13 is performed
        Then a list of crops that grow in 13 days must be returned

    Scenario: A less than time of 5 is specified
        # stardew-crops search --growthlt 5
        When a search by less than growth time of 5 days is performed
        Then a list of crops that grow in 5 days or less must be returned

    Scenario: A greater than time of 5 is specified
        # stardew-crops search --growthgt 5
        When a search by greater than growth time of 5 days is performed
        Then a list of crops that grow in 5 days or more must be returned
