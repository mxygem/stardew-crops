Feature: The Info command

    Info provides data from a store about a requested crop

    Scenario: A matched crop has its data returned
    When info for garlic is requested
    Then the output must match the file "./test_data/crops_raw.json"
