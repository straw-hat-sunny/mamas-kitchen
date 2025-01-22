import { RecipeCard } from "../components/recipeCard/recipeCard";
import { useParams } from "react-router-dom";
import {useEffect, useState} from "react";
import Logo from "@/components/logo/logo";

interface Ingredient {
  item: string;
  quantity: number;
  unit: string;
}

interface Recipe {
  title: string;
  type: "appetizer" | "main course" | "dessert" | "drink";
  ingredients: Ingredient[];
  instructions: string[];
}

const default_recepie: Recipe = {
  title: "Pasta",
  type: "main course",
  ingredients: [
    {
      item: "Pasta",
      quantity: 1,
      unit: "packet",
    },
    {
      item: "Tomato",
      quantity: 2,
      unit: "kg",
    },
  ],
  instructions: ["Boil water", "Add pasta", "Add tomato"],
};

const fetchRecipe = async (id: string) => {
  try {
    const response = await fetch(`/api/recipes/${id}`);
    if (!response.ok) {
      throw new Error("Network response was not ok");
    }
    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error fetching recipe:", error);
    return default_recepie;
  }
}

export default function RecipePage() {
  const { id } = useParams<{ id: string }>();
  const [recipe, setRecipe] = useState<Recipe>(default_recepie);

  useEffect(() => {
    if (id) {
      fetchRecipe(id).then((recipe) => setRecipe(recipe));
    }
  }, [id]);
 
 
  return (
    <div className="app-container" >
      <Logo />
      <RecipeCard
        title={recipe.title}
        type={recipe.type}
        ingredients={recipe.ingredients}
        instructions={recipe.instructions}
      />
    </div>
  );
}
