Feature: Search by Continuous Harvest

    Scenario: A list of crops continuously harvestable
        # stardew-crops search --continuous true
        When a search for continuously harvestable crops is performed
        Then a list of crops that are continuously harvestable are returned

    Scenario: A list of crops that disappear when harvested
        # stardew-crops search --continuous false
        When a search for single harvest crops is performed
        Then a list of crops that are not continuously harvestable are returned
