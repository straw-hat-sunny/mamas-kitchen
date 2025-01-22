

interface IngredientProps {
    item: string;
    quantity: number;
    unit: string;
}
interface IngredientCardProps {
    ingredients: IngredientProps[];
}

export function IngredientCard({ingredients}: IngredientCardProps){

    return (
        <div>
            <h2 style={{ fontSize: '18px', fontWeight: 'bold', textAlign: 'center', textDecoration: 'underline' }}>Ingredients:</h2>
            <ul>
                {ingredients.map((ingredient) => (
                <li key={ingredient.item}>
                - {ingredient.quantity} {ingredient.unit} of {ingredient.item}
                </li>
                ))}
            </ul>
        </div>
    )
};