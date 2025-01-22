
import {Image} from 'semantic-ui-react'
import logo from '../../assets/logo.jpeg'
import { useNavigate } from 'react-router-dom';




export default function Logo() {
    const navigate = useNavigate();
    const handleClick = () => {
        navigate('/recipes')
    }
    return (
    <div onClick={handleClick} style={{ position: "absolute", top: 0, left: 0, display: "flex", alignItems: "center" }}>
        <Image src={logo} size='medium' verticalAlign='top' circular/>
        <h1 style={{ marginLeft: "10px" }}>Mama's Kitchen</h1>
    </div>
    );
    
}