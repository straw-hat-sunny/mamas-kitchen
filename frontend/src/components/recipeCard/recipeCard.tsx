import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "../ui/card";
import { RecipeCarosuel } from "./recipeCarosuel";

interface IngredientProps {
    item: string;
    quantity: number;
    unit: string;
}

interface RecipeProps {
    title: string;
    type: "appetizer" | "main course" | "dessert" | "drink";
    ingredients: IngredientProps[];
    instructions: string[];
}

export function RecipeCard({title, type, ingredients, instructions}: RecipeProps){
    return (
        <Card className="bg-blue-900 text-white" style={{ transform: "scale(1.5)" }}>
            <CardHeader>
            <CardTitle style={{ fontSize: "24px" }}>{title}</CardTitle>
            <CardDescription style={{ fontSize: "12", textDecoration: "underline" }}>{type}</CardDescription>
            </CardHeader>
            <CardContent>
            <RecipeCarosuel ingredients={ingredients} instructions={instructions} />
            </CardContent>
        </Card>
    )
};