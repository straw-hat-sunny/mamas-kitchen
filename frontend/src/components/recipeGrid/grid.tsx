
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "../ui/card";
import { Grid, GridColumn, GridRow } from 'semantic-ui-react';
import { useNavigate } from "react-router-dom";

interface ItemProps {
    id: number;
    title: string;
    type: string;
}



const Item = ({ id, title, type }: ItemProps) => {
    const navigate = useNavigate();
    const handleClick = () => {
        console.log(id);
        navigate(`/recipes/${id}`);
        return
    };

    return (
        <Card className="bg-blue-900 text-white" onClick={handleClick}>
            <CardHeader>
                <CardTitle style={{ fontSize: "24px" }}>{title}</CardTitle>
                <CardDescription style={{ fontSize: "12px", textDecoration: "underline" }}>{type}</CardDescription>
            </CardHeader>
            <CardContent>
                {/* Additional content can go here */}
            </CardContent>
        </Card>
    );
};

interface GridProps {
    items: ItemProps[];
}

const RecipeGrid = ({items}: GridProps) => {
    return (
        <Grid columns={3} divided>
            {items.map((_, index) => (
                index % 3 === 0 && (
                    <GridRow key={index}>
                        <GridColumn length={5} width={5}>
                            <Item id={items[index].id} title={items[index].title} type={items[index].type} />
                        </GridColumn>
                        {items[index + 1] && (
                            <GridColumn width={5}>
                                <Item id={items[index + 1].id} title={items[index + 1].title} type={items[index + 1].type} />
                            </GridColumn>
                        )}
                        {items[index + 2] && (
                            <GridColumn width={5}>
                                <Item id={items[index + 2].id} title={items[index + 2].title} type={items[index + 2].type} />
                            </GridColumn>
                        )}
                    </GridRow>
                )
            ))}
        </Grid>
    )
}

export default RecipeGrid;

