import express, { Request, Response, NextFunction } from "express";
import { supabase } from "../services/supabase.service";
import { BadRequestError, NotFoundError } from "@underthehoodjs/commonjs";

const router = express.Router();

router.post(
  "/api/users/sign-in",
  async (req: Request, res: Response, next: NextFunction) => {
    const { email, password } = req.body;
    try {
      const { data, error } = await supabase.auth.signInWithPassword({
        email: email,
        password: password,
      });
      console.log(error.status);
      if (error) {
        throw new NotFoundError(error.message);
      }
      res.status(200).send(data);
    } catch (error) {
      next(error);
    }
  }
);

export { router as signInRouter };
