import express, { Request, Response, NextFunction } from "express";
import { supabase } from "../services/supabase.service";
import { BadRequestError } from "@underthehoodjs/commonjs";

const router = express.Router();

router.post(
  "/api/users/update-password",
  async (req: Request, res: Response, next: NextFunction) => {
    try {
      const { password } = req.body;

      const { data, error } = await supabase.auth.updateUser({
        password: password,
      });

      if (error) {
        throw new BadRequestError(error.message);
      }

      res.status(200).json(data);
    } catch (error) {
      next(error);
    }
  }
);

export { router as updatePasswordRouter };
