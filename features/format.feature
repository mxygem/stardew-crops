Feature: Formatting options
  Available options:
   * pretty - When outputting data, this will display things in a table or in other nicely formatted ways.
   * json - output formatted JSON
   * raw - (default for now) outputs non-formatted JSON

  Scenario: The pretty formatter returns nicely formatted data
    Given the pretty formatter is desired
    When info for starfruit is requested
    Then the output must match
      """
╔════════════════════════════════════════════════╗
║ Name: Starfruit                                ║
║ Seeds: Starfruit Seeds                         ║
║ Growth Time: 13 days                           ║
║ Season: Summer                                 ║
║ Continual: False                               ║
║ Seed Price:                                    ║
║   * Oasis: 400g                                ║
║ Quests:                                        ║
║   * A Soldier's Star                           ║
║ Notes:                                         ║
║   * Starfruit produces Artisan Goods that have ║
║     some of the highest sell values in the     ║
║     game.                                      ║
║   * Starfruit is required to build a Junimo    ║
║     Hut, purchased from the Wizard's Tower.    ║
╚════════════════════════════════════════════════╝
      """

  Scenario: The json-pretty formatter returns json data with whitespace
    When a search is run with the json formatter
    Then the output must be formatted json

  Scenario: The raw formatter returns search json data with no whitespace
    When a search is run with the raw formatter
    Then the output must be raw json
