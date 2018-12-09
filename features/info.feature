Feature: The Info command

    Info provides data from a store about a requested crop

    Scenario: A matched crop has its data returned
    When info for garlic is requested
    Then the output must match
      """
      {"garlic": {"info": {"description": "Adds a wonderful zestiness to dishes. High quality garlic can be pretty spicy.", "seed": "garlic seeds", "growth_time": 4, "season": "spring", "continual": false}, "seed_price": {"general_store": 40,"jojamart": "N/A", "traveling_cart": {"min": 100, "max": 1000}, "night_market": 40}, "gifting": {"likes": [], "dislikes": [],"hates": []}, "quests": [], "bundles": [], "notes": []}}
      """
