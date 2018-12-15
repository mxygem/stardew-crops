Feature: The Info command

    Info provides data from a store about a requested crop

    Scenario: A crop name must be specified
        # stardew-crops info
        When info is requested
        Then a message informing no crop was specified must be returned

    Scenario: A matched crop has its data returned
        # stardew-crops info garlic
        When info for garlic is requested
        Then the info for garlic must be returned

    Scenario: An unmatched crop name returns a message informing of such
        # stardew-crops info banana
        When info for banana is requested
        Then a message informing no matching crop was found must be returned
