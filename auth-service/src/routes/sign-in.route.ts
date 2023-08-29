import express, { Request, Response, NextFunction } from "express";
import { BadRequestError, NotFoundError } from "@underthehoodjs/commonjs";
import { prisma } from "../services/prisma.service";
import { Password } from "../services/hashing-service";
import jwt from "jsonwebtoken";

const router = express.Router();

router.post(
  "/api/users/sign-in",
  async (req: Request, res: Response, next: NextFunction) => {
    const { email, password } = req.body;
    try {
      const userExists = await prisma.user.findFirst({
        where: {
          email: String(email),
        },
      });

      if (!userExists) {
        throw new NotFoundError("User not found");
      }

      const isValidPassowrd = Password.validatePassowrd(
        userExists.password,
        password
      );

      if (!isValidPassowrd) {
        throw new BadRequestError("Please provide valid credentials");
      }

      // sign jwt with user id and email
      const userJwt = jwt.sign(
        {
          id: userExists.id,
          email: userExists.email,
        },

        process.env.JWT_KEY
      );

      // create a session obj with jwt property and attach it to the request obj
      req.session = {
        jwt: userJwt,
      };
    } catch (error) {
      next(error);
    }
  }
);

export { router as signInRouter };
