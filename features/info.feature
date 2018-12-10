Feature: The Info command

    Info provides data from a store about a requested crop

    Scenario: A crop name must be specified
        When info is requested
        Then the output must match "No crop name specified, please try again"

    Scenario: A matched crop has its data returned
        When info for garlic is requested
        Then the output must match the file "./test_data/crop_garlic_raw.json"

    Scenario: An unmatched crop name returns a message informing of such
        When info for bananas is requested
        Then the output must match "Unable to find matching crop for bananas"
