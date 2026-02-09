package io.github.thunkier.thunkmetrc.client

class MetrcException(
    val statusCode: Int,
    message: String,
    val validationErrors: List<String> = emptyList()
) : RuntimeException(message)
