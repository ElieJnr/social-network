import Image from "next/image";
import Link from "next/link";

export default function NotFound() {
  return (
    <div className="flex  flex-col items-center justify-center min-h-screen  p-6">
      <div className="text-center flex justify-center items-center ">
        <div>
          <h1 className="text-6xl font-bold  mb-4">404</h1>
          <h2 className="text-2xl font-semibold mb-2">Page Not Found</h2>
          <p className="text-gray-600 mb-8">
            The page you are looking for does not exist or has been moved.
          </p>
        </div>
        <div className="mb-8">
          <Image
            src="/error404.jpg"
            width={500}
            height={300}
            alt="Error image"
          />
        </div>
      </div>
        <Link
          href="/"
          className="text-white bg-blue-600 hover:bg-blue-700 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg px-5 py-2.5 text-center"
        >
          Go Back to Home
        </Link>
    </div>
  );
}
