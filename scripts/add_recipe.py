import pathlib
from typing import Optional
from pathlib import Path

import fire
import git

RECIPE_PATH = "recipes"
RECIPE_TEMPLATE: str = """## __name__

### Zutaten

- 


### Zubereitung

1. 

"""


def create_recipe(recipe_name: str, recipe_kind: Optional[str] = None):
    recipe_filename = f"{recipe_name.lower().replace(' ', '_')}.md"
    if recipe_kind:
        p = Path(RECIPE_PATH) / recipe_kind
        Path(p).mkdir(exist_ok=True)
        p = p / recipe_filename
    else:
        p = Path(RECIPE_PATH) / recipe_filename

    recipe = RECIPE_TEMPLATE.replace('__name__', recipe_name)
    with open(p, 'w') as f:
        f.write(recipe)
    repo = git.Repo(pathlib.Path.cwd())
    repo.git.add(p)


if __name__ == "__main__":
    fire.Fire(create_recipe)
