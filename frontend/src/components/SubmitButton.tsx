import React from "react";

type PropsType = {
    buttonName: string;
};

const SubmitButton = ({ buttonName }: PropsType) => {
    return (
        <button
            type="submit"
            data-testid="submit-button"
            className="inline-flex items-center rounded-lg px-5 py-2.5 text-center text-xs font-medium uppercase hover:bg-blue-600 hover:"
        >
            {buttonName}
        </button>
    );
};

export default SubmitButton;
