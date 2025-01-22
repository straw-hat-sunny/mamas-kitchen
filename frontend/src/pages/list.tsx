import FileUpload from "@/components/fileUpload/upload";
import Logo from "@/components/logo/logo";
import RecipeGrid from "@/components/recipeGrid/grid"
import { useEffect, useState } from "react";

const default_items = [
    {
        id: 1,
        title: "Speghetti Bolognese",
        type: "main course" 
    },
    {
        id: 2,
        title: "Pizza",
        type: "main course" 
    },
    {
        id: 3,
        title: "Apple Pie",
        type: "dessert" 
    },
    {  
        id: 4,
        title: "Ice Cream",
        type: "dessert" 
    },
    {
        id: 5,
        title: "Margarita",
        type: "drink" 
    },
    {
        id: 6,
        title: "Mojito",
        type: "drink" 
    },
    {
        id: 7,
        title: "Rum and Coke",
        type: "drink" 
    },
    {
        id: 8,
        title: "Margarita Pizza",
        type: "main course" 
    }
]

const fetchRecipes = async () => {
    try {
        const response = await fetch("/api/recipes/");
        if (!response.ok) {
            throw new Error("Network response was not ok");
        }
        const data = await response.json();
        return data.recipes;
    } catch (error) {
        console.error("Error fetching recipes:", error);
        return default_items;
    }
};


export default function ListPage(){
    const [items, setItems] = useState([]);
    useEffect(() => {
        const getRecipes = async () => {
            const recipes = await fetchRecipes();
            setItems(recipes);
        };
        getRecipes();
    }, []);

    return (
        <div className="list-page">
            <Logo/>
            <FileUpload/>
            <RecipeGrid items={items} />
        </div>
    )
}