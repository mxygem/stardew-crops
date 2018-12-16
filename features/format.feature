@wip
Feature: Formatting options
#   Available options:
#    * Pretty - (default) When outputting data, this will display things in a table or in other nicely formatted ways.
#    * JSON-Pretty - output formatted JSON
#    * Raw - outputs non-formatted JSON

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

#   Scenario: The raw formatter returns json data with no whitespace
#     Given the pretty formatter is desired
#     When info for garlic is requested
#     Then the output must match
#       """
#       "garlic": {"info": {"description": "Adds a wonderful zestiness to dishes. High quality garlic can be pretty spicy.","seed": "garlic seeds","growth_time": 4,"season": "spring","continual": false},"seed_price": {"general_store": 40,"jojamart": "N/A","traveling_cart": {"min": 100,"max": 1000},"night_market": 40}}
#       """
