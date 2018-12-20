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

#   Scenario: The json-pretty formatter returns json data with whitespace
#     Given the pretty formatter is desired
#     When info for garlic is requested
#     Then the output must match
#       """
#       "garlic":{
#           "info": {
#               "description": "Adds a wonderful zestiness to dishes. High quality garlic can be pretty spicy.",
#               "seed": "garlic seeds",
#               "growth_time": 4,
#               "season": "spring",
#               "continual": false
#           },
#           "seed_price": {
#               "general_store": 40,
#               "jojamart": "N/A",
#               "traveling_cart": {
#                   "min": 100,
#                   "max": 1000
#               },
#               "night_market": 40
#           }
#       }
#       """


  Scenario: The json-pretty formatter returns json data with whitespace
    When a search is run with the json formatter
    Then the output must match
      """
      [
        {
          "name": "blueberry",
          "info": {
            "description": "A popular berry reported to have many health benefits. The blue skin has the highest nutrient concentration.",
            "seed": "blueberry seeds",
            "growth_time": 13,
            "season": [
              "summer"
            ],
            "continual": true,
            "regrowth": 4
          },
          "seed_prices": {
            "general_store": 80
          },
          "bundles": [
            "Summer Crops"
          ],
          "recipes": [
            "Blueberry Tart",
            "Fruit Salad"
          ],
          "notes": [
            "Has a small chance for multiple fruit from each harvest"
          ]
        }
      ]
      """

  Scenario: The raw formatter returns info json data with no whitespace
    When info for garlic is requested in raw format
    Then the output must match
      """
      {"name":"garlic","info":{"description":"Adds a wonderful zestiness to dishes. High quality garlic can be pretty spicy.","seed":"garlic seeds","growth_time":4,"season":["spring"],"continual":false},"seed_prices":{"general_store":40},"recipes":["Escargot","Fiddlehead Risotto","Oil of Garlic"],"notes":["Only available starting in year 2"]}
      """

  Scenario: The raw formatter returns search json data with no whitespace
    When a search is run with the raw formatter
    Then the output must match
      """
      [{"name":"blueberry","info":{"description":"A popular berry reported to have many health benefits. The blue skin has the highest nutrient concentration.","seed":"blueberry seeds","growth_time":13,"season":["summer"],"continual":true,"regrowth":4},"seed_prices":{"general_store":80},"bundles":["Summer Crops"],"recipes":["Blueberry Tart","Fruit Salad"],"notes":["Has a small chance for multiple fruit from each harvest"]}]
      """
