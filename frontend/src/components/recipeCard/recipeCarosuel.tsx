import { IngredientCard } from "./ingredientCard";
import { InstructionCard } from "./instructionCard";
import { Carousel, CarouselContent, CarouselItem, CarouselNext, CarouselPrevious } from "../ui/carousel";

interface IngredientProps {
    item: string;
    quantity: number;
    unit: string;
}


interface RecipeCarosuelProps {
    ingredients: IngredientProps[];
    instructions: string[];
}


export function RecipeCarosuel({ingredients, instructions}: RecipeCarosuelProps){
    return (
        <Carousel className="w-full max-w-xs">
            <CarouselContent>
                <CarouselItem>
                    <IngredientCard ingredients={ingredients} />
                </CarouselItem>
                <CarouselItem>
                    <InstructionCard steps={instructions} />
                </CarouselItem>
            </CarouselContent>
            <CarouselPrevious className="text-black" />
            <CarouselNext className="text-black" />
        </Carousel>
    )
}
