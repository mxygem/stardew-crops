Feature: Search by Bundles
    
    Scenario: A specified bundle that uses crops will return the list of crops it uses
        # stardew-crops search --bundle "Summer Crops"
        When a search by bundle for Summer Crops is performed
        Then a list of the crops available in the Summer Crops bundle must be returned

    Scenario: An unmatched bundle will return a not found message
        # stardew-crops search --bundle "Pasta"
        When a search by bundle for Pasta is performed
        Then an error indicating that no match was found must be returned

    Scenario: All crops in bundles are returned when no crop or bundle name is provided
        # stardew-crops search --bundle
        When a search by bundle with no value is performed
        Then an error indicating a value must be specified must be returned
