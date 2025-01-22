

interface InstructionProps {
    steps: string[]; 
}

export function InstructionCard({steps}: InstructionProps){
    return (
        <div>
            <h2 style={{ fontSize: '18px', fontWeight: 'bold', textAlign: 'center', textDecoration: 'underline' }}>Instructions:</h2>
            <ol>
                {steps.map((step, index) => (
                    <li key={index}>
                        {index + 1}. {step}
                    </li>
                ))}
            </ol>
        </div>
    )
};