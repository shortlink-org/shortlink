import React from "react";
const Index = () => {
    return (
        <div className="xl:w-5/12 w-11/12 mx-auto mb-4 my-6 md:w-2/3 shadow sm:px-10 px-4 py-6 bg-white dark:bg-gray-800 rounded-md">
            <p className="text-lg text-gray-800 dark:text-gray-100 font-semibold mb-4">Your Payment Method</p>
            <div className="flex bg-gray-100 dark:bg-gray-700 rounded-md relative">
                <div className="flex">
                    <div className="px-4 py-6 border-r border-gray-200 dark:border-gray-800">
                        <svg width={49} height={38} xmlns="http://www.w3.org/2000/svg" xmlnsXlink="http://www.w3.org/1999/xlink">
                            <image x={55} y={80} width={49} height={38} href="https://tuk-cdn.s3.amazonaws.com/assets/master.png" transform="translate(-55 -80)" fill="none" fillRule="evenodd" />
                        </svg>
                    </div>
                    <div className="flex flex-col justify-center pl-3 text-gray-800 dark:text-gray-100">
                        <p className="text-sm font-bold pb-1">Ending with 4242</p>
                        <div className="flex flex-col sm:flex-row items-start sm:items-center">
                            <p className="text-xs leading-5">Expires 06/21 &nbsp; - &nbsp;</p>
                            <p className="text-xs leading-5">Last updated on 14 March 2020</p>
                        </div>
                    </div>
                </div>
                <div className="w-5 absolute inset-0 m-auto mt-2 sm:mt-4 mr-2 sm:mr-4 sm:right-0 text-indigo-500 cursor-pointer">
                    <svg xmlns="http://www.w3.org/2000/svg" className="icon icon-tabler icon-tabler-edit" width={24} height={24} viewBox="0 0 24 24" strokeWidth="1.5" stroke="currentColor" fill="none" strokeLinecap="round" strokeLinejoin="round">
                        <path stroke="none" d="M0 0h24v24H0z" />
                        <path d="M9 7 h-3a2 2 0 0 0 -2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2 -2v-3" />
                        <path d="M9 15h3l8.5 -8.5a1.5 1.5 0 0 0 -3 -3l-8.5 8.5v3" />
                        <line x1={16} y1={5} x2={19} y2={8} />
                    </svg>
                </div>
            </div>
        </div>
    );
};
export default Index;
