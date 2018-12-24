Feature: Formatting options
  Available options:
   * pretty - (default) When outputting data, this will display things in a table or in other nicely formatted ways.
   * json - output formatted JSON
   * raw - outputs non-formatted JSON

#   @wip
#   Scenario: The pretty formatter returns nicely formatted data
#     Given the pretty formatter is desired
#     When info for garlic is requested
#     Then the output must match
#       """
#       ╔═══════════════════════════════╗
#       ║ Name: Garlic                  ║
#       ║ Seeds: Garlic Seeds           ║
#       ║ Growth Time: 4 days           ║
#       ║ Season: Spring                ║
#       ║ Continual: False              ║
#       ║ Seed Price:                   ║
#       ║   * General Store: 40g        ║
#       ║   * JojaMart: N/A             ║
#       ║   * Traveling Cart: 100-1000g ║
#       ║   * Night Market: 40g         ║
#       ║ Notes: Only avail year 2+     ║
#       ╚═══════════════════════════════╝
#       """

  Scenario: The json-pretty formatter returns json data with whitespace
    When a search is run with the json formatter
    Then the output must be formatted json

  Scenario: The raw formatter returns search json data with no whitespace
    When a search is run with the raw formatter
    Then the output must be raw json
